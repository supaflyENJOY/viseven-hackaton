import React from 'react'

function Exercise(props) {
    const {exercise}=props;
    const musclesList = exercise.Muscles.map(muscle =>
    <li id={muscle.ID}>{muscle.Name}</li>
    );
    const body=<section id={exercise.ID}>
        <img />
        <div className='musclesList_description'>
        {exercise.Title}
        <ul className='musclesList'>
            {musclesList}
        </ul>
        </div>
    </section>;
    return(
        <div>
            {body}
            <p>{exercise.Description}</p>
        </div>
    )
}
export default Exercise;