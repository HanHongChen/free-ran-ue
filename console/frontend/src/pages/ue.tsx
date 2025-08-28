import { Outlet } from 'react-router-dom'
import styles from './css/ue.module.css'
import Sidebar from '../components/sidebar/sidebar'
import { useUe } from '../context/ueContext'

export default function Ue() {
  const { ranUeList, xnUeList } = useUe()

  return (
    <div className={styles.container}>
      <Sidebar />
      <div className={styles.content}>
        <div className={styles.header}>
          <h1>UE</h1>
        </div>

        <div className={styles.infoCard}>
          <h2 className={styles.title}>RAN UE List</h2>
          <div className={styles.ueList}>
            <table className={styles.table}>
              <thead className={styles.tableHeader}>
                <tr>
                  <th>No.</th>
                  <th>UE</th>
                  <th>gNB</th>
                  <th>DC-status</th>
                </tr>
              </thead>
              <tbody>
                {ranUeList.map((ue, index) => (
                  <tr key={ue.imsi}>
                    <td>{index + 1}</td>
                    <td>{ue.imsi}</td>
                    <td>{ue.gnbName || ue.gnbId}</td>
                    <td>
                      <span className={ue.nrdcIndicator ? styles.statusOnline : styles.statusOffline}>
                        {ue.nrdcIndicator ? 'Enabled' : 'Disabled'}
                      </span>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>

        <div className={styles.infoCard}>
          <h2 className={styles.title}>XN UE List</h2>
          <div className={styles.ueList}>
            <table className={styles.table}>
              <thead className={styles.tableHeader}>
                <tr>
                  <th>No.</th>
                  <th>UE</th>
                  <th>gNB</th>
                </tr>
              </thead>
              <tbody>
                {xnUeList.map((ue, index) => (
                  <tr key={ue.imsi}>
                    <td>{index + 1}</td>
                    <td>{ue.imsi}</td>
                    <td>{ue.gnbName || ue.gnbId}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
        <Outlet />
      </div>
    </div>
  )
}
