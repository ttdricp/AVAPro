import React from 'react'
import {Container} from "react-bootstrap";
import './CoursePage.css';
import GridView  from '../components/grid-view/GridView';
import {useDispatch, useSelector} from "react-redux";

const TasksPage = () => {
    // const dispatch = useDispatch()
    // const  tasks = useSelector(state=>state.client.tasks)
    const tasks = [
        {
            title: "Arakasi",
            description: "somedesctiprion",

        }

    ]
    // useEffect(()=>{
    //     dispatch(fetchTasks())
    // }, [])
    let key = useSelector(state => state.client.key)
    let sort = useSelector(state => state.client.sort)
    let currentPage = useSelector(state =>state.client.page)
    let total = useSelector(state=>state.client.total)
    return (
        <Container className={'py-5'}>

                <GridView items={tasks}>
                    <div className={'d-flex flex-row justify-content-between'}>
                        <h1 className={'p-0 m-0'}>Tasks</h1>
                        {/* <Button className={''} onClick={()=>dispatch(fetchCourses('title', 'asc',1 ))}>Sort by Title</Button>
                        <Button className={''} onClick={()=>dispatch(fetchCourses('category', 'asc',1 ))}>Sort by Category</Button>
                        <ButtonGroup className={''}>
                            <Button disabled={currentPage <= 1} onClick={()=>dispatch(fetchCourses(key, sort, --currentPage))}>Previous Page</Button>
                            <Button disabled>{currentPage}</Button>
                            <Button disabled={currentPage >= total} onClick={()=>dispatch(fetchCourses(key, sort, ++currentPage))}>Next Page</Button>
                        </ButtonGroup> */}
                    </div>
                </GridView>

                
        </Container>
    );
};



export default TasksPage;