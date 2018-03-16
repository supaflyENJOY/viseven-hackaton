import React from 'react';
import Muscle from '../Muscle'

export default function ({muscles}) {
    const musclesElements=muscles.map(muscle =>
        <li key={muscle.id}><Muscle muscle={muscle}/></li>
    );
    return(
        <ul>
            {musclesElements}
        </ul>
    )
}
