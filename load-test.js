import http from 'k6/http';
import { sleep } from 'k6';
export const options = {
    vus: 10,
    duration: '100s',
};

const host = 'http://localhost:9090';
export default function () {

    GetList()
    Login()
    Register()
    ForgotPassword()
    Update()

}

function Login(){
    const url = host+'/login';
    let id = Math.floor((Math.random() * 100) + 1)
    const payload = JSON.stringify({
        email: id.toString(),
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
    let id = Math.floor((Math.random() * 100) + 1)
    const payload = JSON.stringify({
        email: 'email'+id,
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
    let id = Math.floor((Math.random() * 100) + 1)
    const payload = JSON.stringify({
        name: 'name'+id,
        email: id.toString(),
        password: 'password'+id,
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.post(url, payload, params);
}

function Update(){
    const url = host+'/user';
    let id = Math.floor((Math.random() * 100) + 1)
    const payload = JSON.stringify({
        name: 'name'+id,
        email: id.toString(),
        password: 'password'+id,
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.put(url, payload, params);
}

function GetList(){
    const url = host+'/users';

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.get(url,  params);
}