import { AuthenticationApi } from './api'
import axios from 'axios'

const getBaseUrl = () => {
    const { protocol, hostname } = window.location
    return `${protocol}//${hostname}:40104`
}

const apiConfig = {
    basePath: getBaseUrl(),
    isJsonMime: () => false,
}

export const authApi = new AuthenticationApi(apiConfig, undefined, axios)