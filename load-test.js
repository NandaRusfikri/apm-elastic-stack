import http from 'k6/http';
import { sleep } from 'k6';
export const options = {
    vus: 10,
    duration: '10s',
};

const host = 'http://localhost:9090';
export default function () {

    Login()
    Register()
    ForgotPassword()
}

function Login(){
    const url = host+'/login';
    const payload = JSON.stringify({
        email: 'email',
        password: 'password',
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.post(url, payload, params);
}

function ForgotPassword(){
    const url = host+'/forgot-password';
    const payload = JSON.stringify({
        email: 'email',
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.post(url, payload, params);
}

function Register(){
    const url = host+'/register';
    const payload = JSON.stringify({
        name: 'name',
        email: 'email',
        password: 'password',
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.post(url, payload, params);
}