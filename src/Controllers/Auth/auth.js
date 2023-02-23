import AuthService from "../../Services/AuthService"
import {setUser, logout, setTask, setTasks, setUserTasks} from "../../Reducers/ClientReducer";

// import {states} from "../reducers/adminReducer";


export const logoutPage = () => {
    return async dispatch => {
        try{
            const response = await AuthService.logout()
            dispatch(logout())
        }catch (e){ 
            alert(e.response.data.message)
         }
    }
}

export const login = (email, password) => {
    return async dispatch => {
        try{
            const response = await AuthService.login(email, password)
            localStorage.setItem('token', response.data.accessToken)
            dispatch(setUser(response.data.user))
        }catch (e) {
            alert(e.response.data.message)
        }
    }
}

export const registration = (name, email, password) => {
    return async dispatch => {
        try{
            const response = await AuthService.registration(name, email, password)
            localStorage.setItem('token', response.data.accessToken)
            dispatch(setUser(response.data.user))
        }catch (e) {
            alert(e.response.data.message)
        }
    }
}


