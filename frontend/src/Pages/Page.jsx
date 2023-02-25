import React from 'react'
import {Outlet} from "react-router-dom"
import {Container} from "react-bootstrap";
import Navigation from '../components/navigation/navigation';
import Footer from '../components/footer/footer';


const Page = () => {
    return (
        <>
            <Container fluid className="p-0">
                <div >
                    <Navigation/>
                    <div className='d-flex min-vh-100'>
                        <Outlet />
                    </div>
                    <Container>
                    <Footer/>
                    </Container>
                </div>
            </Container>
        </>
    )
}

export default Page;