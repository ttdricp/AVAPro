import React from 'react';
import {Col, Modal, Row} from "react-bootstrap";
import {Link} from "react-router-dom";
import {useDispatch} from "react-redux";

const GridItemModal = (props) => {
    const dispatch= useDispatch()
    return (
        <Modal
            {...props}
            size="xl"
            aria-labelledby="contained-modal-title-vcenter"
            centered
        >
            <Modal.Header closeButton>
                <Modal.Title id="contained-modal-title-vcenter">
                    {props.item.title}
                </Modal.Title>
                <Link to={`../tasks/${props.item.id}`} className={"btn ms-auto btn-maximize"} 

                // TODO: фетч курсов
                // onClick={()=>dispatch(fetchTasks(props.item.id))}
                />

            </Modal.Header>
            <Modal.Body as={Row} className={'row-cols-1 row-cols-lg-2'}>
                <Col className={'d-flex flex-column'}>
                    <div className={'flex-grow-1'}>
                        <h5>Description</h5>
                        <p>{props.item.description}</p>
                        <p>{props.item.points}</p>
                    </div>
                </Col>
            </Modal.Body>
        </Modal>
    );
};

export default GridItemModal;