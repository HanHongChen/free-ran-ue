#! /bin/bash

FREE_MRAN_NAMESPACE="free-mran-ns"
FREE_SRAN_NAMESPACE="free-sran-ns"
FREE_UE_NAMESPACE="free-ue-ns"

IPERF_A_NAMESPACE="iperf-a-ns"
IPERF_B_NAMESPACE="iperf-b-ns"

usage() {
    echo "Usage: $0 [ up | down | mran-ns | sran-ns | ue-ns | iperf-a | iperf-b ]"
    echo "  up     - Setup both MRAN and SRAN namespaces"
    echo "  down   - Cleanup both namespaces"
    echo "  mran-ns - Enter Master RAN namespace"
    echo "  sran-ns - Enter Secondary RAN namespace"
    echo "  ue-ns  - Enter UE namespace"
    echo "  iperf-a - Enter iperf-a namespace"
    echo "  iperf-b - Enter iperf-b namespace"
    exit 1
}

setup_network_namespace() {
    # Remove exist network namespace
    echo "Removing exist network namespace..."
    cleanup_network_namespace
    echo

    # Create network namespace
    echo "Creating network namespace..."
    sudo ip netns add $FREE_MRAN_NAMESPACE 2>/dev/null || true
    sudo ip netns add $FREE_SRAN_NAMESPACE 2>/dev/null || true
    sudo ip netns add $FREE_UE_NAMESPACE 2>/dev/null || true
    sudo ip netns add $IPERF_A_NAMESPACE 2>/dev/null || true
    sudo ip netns add $IPERF_B_NAMESPACE 2>/dev/null || true
    echo

    # Create veth pair
    echo "Creating veth pair and bridge..."
    # Create bridge for connecting MRAN and SRAN to same host interface
    sudo ip link add brHost type bridge
    sudo ip link add br-mran type veth peer fnsMRAN
    sudo ip link add br-sran type veth peer fnsSRAN
    sudo ip link add fnsMGnb type veth peer fnsMUe
    sudo ip link add fnsSGnb type veth peer fnsSUe

    sudo ip link add brIperf type bridge
    sudo ip link add br-iperf-a type veth peer fnsIperfA
    sudo ip link add br-iperf-b type veth peer fnsIperfB
    
    # Connect veth pairs to bridge
    sudo ip link set br-mran master brHost
    sudo ip link set br-sran master brHost

    sudo ip link set br-iperf-a master brIperf
    sudo ip link set br-iperf-b master brIperf
    echo

    # Move veth pair to network namespace
    echo "Moving veth pair to network namespace..."
    sudo ip link set fnsMRAN netns $FREE_MRAN_NAMESPACE
    sudo ip link set fnsSRAN netns $FREE_SRAN_NAMESPACE
    sudo ip link set fnsMGnb netns $FREE_MRAN_NAMESPACE
    sudo ip link set fnsSGnb netns $FREE_SRAN_NAMESPACE
    sudo ip link set fnsMUe netns $FREE_UE_NAMESPACE
    sudo ip link set fnsSUe netns $FREE_UE_NAMESPACE

    sudo ip link set fnsIperfA netns $IPERF_A_NAMESPACE
    sudo ip link set fnsIperfB netns $IPERF_B_NAMESPACE
    echo

    # Bring up the interface in host namespace
    echo "Bring up the interface in host namespace..."
    sudo ip link set brHost up
    sudo ip link set br-mran up
    sudo ip link set br-sran up

    sudo ip link set brIperf up
    sudo ip link set br-iperf-a up
    sudo ip link set br-iperf-b up
    echo

    # Bring up the interface in free-ran-ue namespace
    echo "Bring up the interface in free-ran-ue namespace..."
    sudo ip netns exec $FREE_MRAN_NAMESPACE ip link set fnsMRAN up
    sudo ip netns exec $FREE_SRAN_NAMESPACE ip link set fnsSRAN up
    sudo ip netns exec $FREE_MRAN_NAMESPACE ip link set fnsMGnb up
    sudo ip netns exec $FREE_SRAN_NAMESPACE ip link set fnsSGnb up
    echo

    # Bring up the interface in free-ran-ue namespace
    echo "Bring up the interface in free-ran-ue namespace..."
    sudo ip netns exec $FREE_UE_NAMESPACE ip link set fnsMUe up
    sudo ip netns exec $FREE_UE_NAMESPACE ip link set fnsSUe up
    echo

    # Bring up the interface in iperf-a namespace
    echo "Bring up the interface in iperf-a namespace..."
    sudo ip netns exec $IPERF_A_NAMESPACE ip link set fnsIperfA up
    echo
    # Bring up the interface in iperf-b namespace
    echo "Bring up the interface in iperf-b namespace..."
    sudo ip netns exec $IPERF_B_NAMESPACE ip link set fnsIperfB up
    echo

    # Bring up the interface in iperf-b namespace
    # Set up IP address
    echo "Setting up IP address..."
    # free-ran-ue namespace: 10.0.1.0/24 network
    sudo ip addr add 10.0.1.1/24 dev brHost
    sudo ip netns exec $FREE_MRAN_NAMESPACE ip addr add 10.0.1.2/24 dev fnsMRAN
    sudo ip netns exec $FREE_SRAN_NAMESPACE ip addr add 10.0.1.3/24 dev fnsSRAN
    sudo ip netns exec $FREE_MRAN_NAMESPACE ip addr add 10.0.2.1/24 dev fnsMGnb
    sudo ip netns exec $FREE_SRAN_NAMESPACE ip addr add 10.0.3.1/24 dev fnsSGnb
    sudo ip netns exec $FREE_UE_NAMESPACE ip addr add 10.0.2.2/24 dev fnsMUe
    sudo ip netns exec $FREE_UE_NAMESPACE ip addr add 10.0.3.2/24 dev fnsSUe

    sudo ip addr add 10.1.0.1/24 dev brIperf
    sudo ip netns exec $IPERF_A_NAMESPACE ip addr add 10.1.0.2/24 dev fnsIperfA
    sudo ip netns exec $IPERF_B_NAMESPACE ip addr add 10.1.0.3/24 dev fnsIperfB
    echo

    # Set up default route
    echo "Setting up default route..."
    sudo ip netns exec $FREE_MRAN_NAMESPACE ip route add default via 10.0.1.1
    sudo ip netns exec $FREE_SRAN_NAMESPACE ip route add default via 10.0.1.1
    sudo ip netns exec $FREE_MRAN_NAMESPACE ip route add 10.0.2.0/24 dev fnsMGnb
    sudo ip netns exec $FREE_SRAN_NAMESPACE ip route add 10.0.3.0/24 dev fnsSGnb
    sudo ip netns exec $FREE_UE_NAMESPACE ip route add 10.0.3.0/24 dev fnsSUe

    sudo ip netns exec $IPERF_A_NAMESPACE ip route add default via 10.1.0.1
    sudo ip netns exec $IPERF_B_NAMESPACE ip route add default via 10.1.0.1
    echo

    echo "$FREE_MRAN_NAMESPACE namespace setup complete"
    echo "$FREE_SRAN_NAMESPACE namespace setup complete"
    echo "$FREE_UE_NAMESPACE namespace setup complete"
    echo "$IPERF_A_NAMESPACE namespace setup complete"
    echo "$IPERF_B_NAMESPACE namespace setup complete"
    echo "Network topology:"
    echo "  Host brHost (10.0.1.1) <---> MRAN namespace (10.0.1.2 | 10.0.2.1) <---> UE namespace (10.0.2.2)"
    echo "  Host brHost (10.0.1.1) <---> SRAN namespace (10.0.1.3 | 10.0.3.1) <---> UE namespace (10.0.3.2)"
    echo "  Host brIperf (10.1.0.1) <---> iperf-a namespace (10.1.0.2)"
    echo "  Host brIperf (10.1.0.1) <---> iperf-b namespace (10.1.0.3)"
    echo "  Note: Both MRAN and SRAN connect to Host via the same bridge interface (brHost) and iperf-a and iperf-b connect to Host via the same bridge interface (brIperf)"
}

cleanup_network_namespace() {
    echo "Removing network namespace..."

    # Bring down interface
    sudo ip link set brHost down 2>/dev/null || true
    sudo ip link set br-mran down 2>/dev/null || true
    sudo ip link set br-sran down 2>/dev/null || true

    sudo ip link set brIperf down 2>/dev/null || true
    sudo ip link set br-iperf-a down 2>/dev/null || true
    sudo ip link set br-iperf-b down 2>/dev/null || true

    # Delete veth pair (deleting one end automatically deletes the pair)
    sudo ip link delete br-mran 2>/dev/null || true
    sudo ip link delete br-sran 2>/dev/null || true
    sudo ip link delete fnsMGnb 2>/dev/null || true
    sudo ip link delete fnsSGnb 2>/dev/null || true

    sudo ip link delete br-iperf-a 2>/dev/null || true
    sudo ip link delete br-iperf-b 2>/dev/null || true
    
    # Delete bridge
    sudo ip link delete brHost 2>/dev/null || true

    sudo ip link delete brIperf 2>/dev/null || true

    # Delete network namespace
    sudo ip netns delete $FREE_MRAN_NAMESPACE 2>/dev/null || true
    sudo ip netns delete $FREE_SRAN_NAMESPACE 2>/dev/null || true
    sudo ip netns delete $FREE_UE_NAMESPACE 2>/dev/null || true

    sudo ip netns delete $IPERF_A_NAMESPACE 2>/dev/null || true
    sudo ip netns delete $IPERF_B_NAMESPACE 2>/dev/null || true

    echo "$FREE_MRAN_NAMESPACE namespace removed"
    echo "$FREE_SRAN_NAMESPACE namespace removed"
    echo "$FREE_UE_NAMESPACE namespace removed"

    echo "$IPERF_A_NAMESPACE namespace removed"
    echo "$IPERF_B_NAMESPACE namespace removed"
}

main() {
    if [ $# -ne 1 ]; then
        usage
    fi

    case "$1" in
        "up")
            setup_network_namespace
        ;;
        "down")
            cleanup_network_namespace
        ;;
        "mran-ns")
            sudo ip netns exec $FREE_MRAN_NAMESPACE bash
        ;;
        "sran-ns")
            sudo ip netns exec $FREE_SRAN_NAMESPACE bash
        ;;
        "ue-ns")
            sudo ip netns exec $FREE_UE_NAMESPACE bash
        ;;
        "iperf-a")
            sudo ip netns exec $IPERF_A_NAMESPACE bash
        ;;
        "iperf-b")
            sudo ip netns exec $IPERF_B_NAMESPACE bash
        ;;
        *)
            usage
        ;;
    esac
}

main "$@"