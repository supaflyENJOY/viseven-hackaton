import React, { Component } from 'react';
import './index.css'
import exercise from '../api/exercise';
import Navbar from '../Navbar';

class MainPart extends Component {
    render(){
        return (<div>
            <Navbar/>
            <div className='exerciseListBox'>
                <ul>
                    {exercise.map(exercise=>
                        <li key={exercise.ID}>


                            <div>
                                {<section id={exercise.ID}>
                                    <img className='imageExercice' />
                                    <a className='titleBox'> {exercise.Title}</a>
                                    <a className='addExercise' href='#'><img src={require('../img/plus.png')}/></a>
                                    <a className='showDetails' href='#' ><img src={require('../img/menu-down.png')}/></a>
                                </section>}
                                <div className='details'>
                                    <p>{exercise.Description}</p>
                                    <ul className='musclesList'>
                                        {exercise.Muscles.map(muscle =>
                                            <li id={muscle.ID}>{muscle.Name}</li>
                                        )}
                                    </ul>
                                </div>
                            </div>

                        </li>
                    )}
                </ul>
            </div>
            </div>
            );
    }
}

export default MainPart;
