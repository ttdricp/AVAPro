import $api from "../Http";
import {AxiosResponse} from "axios";

export default class AuthService{
    static async registration(name, email,  password){
        return $api.post('/auth/registration', {name, email, password})
    }
    static async login(email, password){
        return $api.post('/auth/login', {email, password})
    }
    static async logout(){
        return $api.post('/auth/logout')
    }
    static async refresh(){
        return $api.get('/auth/refresh')
    }
}