import React from 'react';
import Exercise from '../Exercise'


export default function ({exercise}) {
    const exerciseElements=exercise.map(exercise=>
        <li key={exercise.ID}><Exercise exercise={exercise}/></li>
    );
    return(
        <ul>
            {exerciseElements}
        </ul>
    )
}
