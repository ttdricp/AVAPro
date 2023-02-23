import { Routes, Route } from "react-router-dom";

import {useDispatch, useSelector} from "react-redux";

import HomePage from './Pages/HomePage';
import Page from './Pages/Page';


import RankingsPage from './Pages/RankingsPage';
import TasksPage from './Pages/TasksPage';
import SignUpPage from "./Pages/SignUpPage";
import LoginPage from "./Pages/LoginPage";

function App(){
  // clientReducer берет
  const isAuth = useSelector(state => state.client.isAuth);

//   useEffect(() => {
//     (
//       async () => {
//           const response = await fetch('http://localhost:8000/api/user', {
//               headers: {'Content-Type': 'application/json'},
//               credentials: 'include',
//           });

//           const content = await response.json();

//           setName(content.name);
//       }
//   )();
// });

  return <>
  <Routes>

    {/* Если не авторизован, доступ ко всем path */}

  
    <Route path="/" element={<Page/>}>  

    <Route index element={<HomePage/>} />  
      <Route path="tasks" element={<TasksPage/>} />  
      <Route path="rankings" element={<RankingsPage/>} />

      <Route path="signup" element={<SignUpPage/>} />
      <Route path="login" element={<LoginPage/>} />



      
    {/* {
      isAuth ? (
        <Route path="profile/:id" element={<ProfilePage/>} />
      ) : (
        <Route path="auth" element={<AuthPage/>} />
        )
    } */}

    </Route>

  </Routes>
  </>
  
  }
  
  export default App;