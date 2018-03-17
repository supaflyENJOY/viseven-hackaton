import React, { Component } from 'react';
import './index.css'
import Navbar from '../Navbar';
import exercise from "../api/exercise";

class MainPart extends Component {
    constructor(props) {
        super(props);
        this.state = {
            currentShow: -1,
            exercises: [],
            usedMuscles: [],
            templates:[[]]
        };
        this.handleClick = this.handleClick.bind(this);
        fetch('http://petrosyan.in:8080/v1/exercise/find', {
            method: 'POST',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({input: this.state.usedMuscles}),
            credentials: 'include'
        }).then((response) => response.text())
            .then((responseJson) => {
                console.log(responseJson)
            })
            .catch((error) => {
                console.error(error);
            });
    }


    handleClick(id) {
        if(id == this.state.currentShow) {
            this.setState({ currentShow: -1 });
        } else {
            this.setState({ currentShow: id });
        }

    }

    render(){
        return (<div>
            <Navbar/>
            <div className='exerciseListBox'>
                <ul>
                    {this.state.exercises.map(exercise=>
                        <li key={exercise.ID}>
                            <div>
                                {<section id={exercise.ID}>
                                    <img className='imageExercice' />
                                    <a className='titleBox'> {exercise.Title}</a>
                                    <a className='addExercise' href='#'><img src={require('../img/plus.png')}/></a>
                                    <a className='showDetails' href='#' onClick={() => { this.handleClick(exercise.ID) }}><img src={require('../img/menu-down.png')} style={this.state.currentShow == exercise.ID ? {transform: "rotate(180deg)"}: null}/></a>
                                </section>}
                                {this.state.currentShow == exercise.ID?
                                <div className='details'>
                                    <p>{exercise.Description}</p>
                                    <ul className='musclesList'>
                                        {exercise.Muscles.map(muscle =>
                                            <li id={muscle.ID}>{muscle.Name}</li>
                                        )}
                                    </ul>
                                </div>
                                    : null }
                            </div>

                        </li>
                    )}
                </ul>
            </div>
            <div className='templateListBox'>{this.state.templates!==[[]]?
                <a><img className='getNewTemplate' src={require('../img/bigPlus.png')}/></a>:
                <ul className='exercisesList_ForTemplate'>
                    {this.state.templates.map(template =>
                        <ul>{template.map(exercise=>
                            <li>
                                <img className='imageExercice'/>
                                <a className='titleBox'> {exercise.Title}</a>
                                <a className='addExercise' href='#'><img src={require('../img/plus.png')}/></a>
                            </li>
                        )}</ul>
                    )}
                    </ul>
            }


            </div>
            </div>
            );
    }
}

export default MainPart;
