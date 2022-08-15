import axios, {AxiosInstance} from 'axios';
import {ApiBase} from "../config/api";

const ApiClient: AxiosInstance = axios.create({
    baseURL: ApiBase,
    timeout: 15
});

export default ApiClient;