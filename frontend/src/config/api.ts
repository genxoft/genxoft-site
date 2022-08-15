let apiBase;
if (process.env.REACT_APP_VERSION == undefined || process.env.REACT_APP_VERSION === 'dev') {
    apiBase = 'http://localhost:8080/api';
} else {
    apiBase = '/api';
}

export const ApiBase = apiBase;