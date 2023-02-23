import $api from "../Http";
import {AxiosResponse} from "axios";

export default class AuthService{
    // на будущее
    static async rankings(){
        return $api.get('/user/ranking')
    }}
