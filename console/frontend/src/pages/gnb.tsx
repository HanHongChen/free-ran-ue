import { Outlet } from 'react-router-dom'
import styles from './css/gnb.module.css'
import Sidebar from '../components/sidebar/sidebar'

export default function Gnb() {
  return (
    <div className={styles.container}>
      <Sidebar />
      <div className={styles.content}>
        <div className={styles.header}>
          <h1>gNB</h1>
        </div>
        <Outlet />
      </div>
    </div>
  )
}
