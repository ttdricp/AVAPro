import React from 'react';
import GridItem from "./GridItem";

const GridView = (props) => {
    if(props.items?.length === 0){
        return(
            <>
                {props.children}
                <div className={'text-center display-6'}>Not found</div>
            </>
        )
    }
    return (
        <>
            {props.children}
            <div className={'py-5'}>
                <div className={"row row-cols-3 g-2"}>
                    {props?.items && props?.items.map((item, index)=>
                        <GridItem index={index} key={item.title} item={item}/>
                    )}
                </div>
            </div>
        </>
    );
};

export default GridView;