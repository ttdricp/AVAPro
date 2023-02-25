import AuthService from "../../../../../react-frontend/src/Services/AuthService"
import {setUser, logout,setTask, setTasks, setUserTasks, rankings} from "../../Reducers/ClientReducer";

// import {states} from "../reducers/adminReducer";


export const FetchRankings = () => {
    return async dispatch => {
        try{
            const response = [
                { rank: 1, name: 'John Smith', points: 1000, language: 'Java' },
                { rank: 2, name: 'Jane Doe', points: 800, language: 'Java' },
                { rank: 3, name: 'Bob Johnson', points: 500, language: 'Java' },
                { rank: 4, name: 'Samantha Brown', points: 300, language: 'Java' },
                { rank: 5, name: 'Chris Rakhat', points: 200, language: 'Java' },
              ]
            //   const response = await UserService.rankings();
            dispatch(rankings(response))
        }catch (e){ 
            alert(e.response.data.message)
         }
    }
}
