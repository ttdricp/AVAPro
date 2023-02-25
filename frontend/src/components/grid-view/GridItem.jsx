import React, {useState} from 'react';
import {Button} from "react-bootstrap";
import GridItemModal from "./GridItemModal";

const GridItem = (props) => {
    // let colors=['#3FCBFF', '#fed030', '#A448F8']
    const [modalShow, setModalShow] = useState(false);
    return (
        <div className={"col"} style={{height: '20rem'}}>
            <div className={'h-100 rounded-3'} 
            // style={{backgroundColor: colors[props.index%3]}}
            >
                {props.item?.img && (
                    <img src={props.item?.src} alt={"course"}/>
                )}
                <div className={'d-flex flex-column align-items-start h-100 p-3'}>
                    <h3 className={'m-0 p-0 text-decoration-underline'}>{props.item?.title}</h3>
                    <p className={'m-0 p-0 flex-grow-1'}>{props.item?.description}</p>
                    <p className={'m-0 p-0 flex-grow-1'}>{props.item?.points}</p>
                    <div className={'w-100'}>
                        <Button variant={'outline-secondary'}
                                className={'w-100'}
                                onClick={() => setModalShow(true)}>
                            More
                        </Button>
                        <GridItemModal
                            item={props.item}
                            show={modalShow}
                            onHide={() => setModalShow(false)}
                        />
                    </div>
                </div>
            </div>
        </div>
    );
};
export default GridItem;