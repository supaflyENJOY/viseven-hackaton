import React from 'react'

function Muscle(props) {
    const {muscle}=props;
    const body=<section>{muscle.name}</section>;
    return(
        <div>
            {body}
        </div>
    )
}
export default Muscle;