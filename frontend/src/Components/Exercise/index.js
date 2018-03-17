import React  from 'react'

function Exercise(props) {
    const {exercise}=props;
    const musclesList = exercise.Muscles.map(muscle =>
    <li id={muscle.ID}>{muscle.Name}</li>
    );
    const body=<section id={exercise.ID}>
        <img className='imageExercice' />
        <a className='titleBox'> {exercise.Title}</a>
        <a className='addExercise' href='#'><img src={require('../img/plus.png')}/></a>
        <a className='showDetails' href='#' ><img src={require('../img/menu-down.png')}/></a>
    </section>;
    return(
        <div>
            {body}
            <div className='details'>
            <p>{exercise.Description}</p>
            <ul className='musclesList'>
                {musclesList}
            </ul>
            </div>
        </div>
    )
}
export default Exercise;