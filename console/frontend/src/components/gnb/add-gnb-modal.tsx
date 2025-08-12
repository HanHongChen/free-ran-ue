import { useState } from 'react'
import Modal from '../modal/modal'
import styles from '../modal/modal.module.css'
import { useErrors } from '../../hooks/useErrors'
import ErrorBox from '../errorBox/errorBox'
import { gnbApi } from '../../apiCfg'
import { useGnb } from '../../context/gnbContext'

interface AddGnbModalProps {
  isOpen: boolean
  onClose: () => void
}

export default function AddGnbModal({ isOpen, onClose }: AddGnbModalProps) {
  const { errors, addError, removeError } = useErrors()
  const { addGnb } = useGnb()
  const [formData, setFormData] = useState({
    ip: '',
    port: ''
  })

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target
    setFormData(prev => ({
      ...prev,
      [name]: value
    }))
  }

  const handleSubmit = async () => {
    try {
      if (!formData.ip || !formData.port) {
        addError('Please fill in all fields')
        return
      }

      const portNumber = parseInt(formData.port)
      if (isNaN(portNumber) || portNumber < 1 || portNumber > 65535) {
        addError('Invalid port number')
        return
      }

      const result = await gnbApi.apiConsoleGnbRegistrationPost({
        ip: formData.ip,
        port: portNumber
      }, {
        headers: {
          'Authorization': localStorage.getItem('token')
        }
      })

      addGnb(result.data)
      onClose()
    } catch (error: any) {
      addError(error.response?.data?.message || 'Failed to add gNB')
    }
  }

  return (
    <>
      <ErrorBox 
        errors={errors}
        onClose={removeError}
        duration={5000}
      />
      <Modal
      isOpen={isOpen}
      onClose={onClose}
      title="Add gNB"
      onSubmit={handleSubmit}
      >
      <div className={styles.formGroup}>
        <label className={styles.label} htmlFor="ip">
          IP Address
        </label>
        <input
          id="ip"
          name="ip"
          type="text"
          className={styles.input}
          value={formData.ip}
          onChange={handleChange}
          placeholder="Enter IP address"
        />
      </div>

      <div className={styles.formGroup}>
        <label className={styles.label} htmlFor="port">
          Port
        </label>
        <input
          id="port"
          name="port"
          type="text"
          className={styles.input}
          value={formData.port}
          onChange={handleChange}
          placeholder="Enter port number"
        />
      </div>
    </Modal>
    </>
  )
}
