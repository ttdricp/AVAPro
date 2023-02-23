const SET_USER = "SET_USER"
const SET_USERTASKS = "SET_USERTASKS"
const SET_TASKS = "SET_TASKS"
const STATES = "STATES"
const LOGOUT = "LOGOUT"
const isLOADING = "isLOADING"
const SET_RANKS = "SET_RANKS"

const defaultState = {
    user: {},
    isAuth: false,
    isAdmin: false,
    isLoading: false,
    tasks: [],
    usertasks: [],
    ranking: [],
    key: 'title',
    sort: 'asc',
    page: 1,
    total: 1,
}

export default function ClientReducer(state = defaultState, action) {
    switch (action.type){
        case SET_USER:
            return {
                ...state,
                user: action.payload,
                isAdmin: action.payload.roles.includes('admin'),
                isAuth: true,
            }
        case SET_RANKS:
            return{
                ...state,
                ranking: action.payload,
                
            }
        
        case SET_USERTASKS:
            return {
                ...state,
                usertasks: action.payload.tasks,
                page: action.payload.currentPage,
                total: action.payload.totalPages
            }
        case SET_TASKS:
            return {
                ...state,
                tasks: action.payload.tasks,
                page: action.payload.currentPage,
                total: action.payload.totalPages
            }
        case STATES:
            return {
                ...state,
                key: action.payload.key,
                sort: action.payload.sort,
            }
        case LOGOUT:
            localStorage.removeItem('token')
            return {
                ...state,
                user: {},
                isAuth: false,
                isAdmin: false
            }
        case isLOADING:
            return{
                ...state,
                isLoading: action.payload
            }
        default:
            return state
    }
}

export const setUser =  user => ({type: SET_USER, payload: user})
export const setUserTasks = tasks => ({type: SET_USERTASKS, payload: tasks})
export const setTasks =  tasks => ({type: SET_TASKS, payload: tasks})
export const states = obj => ({type: STATES, payload: obj})
export const isLoading =  bool => ({type: isLOADING, payload: bool})
export const logout = () => ({type: LOGOUT})
export const rankings = ranks => ({type: SET_RANKS, payload: ranks})