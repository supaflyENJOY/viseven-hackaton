import React, { Component } from 'react';
import './index.css'
import Navbar from '../Navbar';
import exercise from "../api/exercise";

import { RegisterExternalListener } from "react-unity-webgl";
import Unity from "react-unity-webgl";

class MainPart extends Component {
    constructor(props) {
        super(props);
        this.state = {
            currentShow: -1,
            currentActiveTemplate: -1,
            exercises: [],
            usedMuscles: [],
            templates: [],
            profile: [],
            value: {
                age:'',
                height:'',
                weight:'',
                phone:''
            }

        };

        RegisterExternalListener("updateSelectedMuscles", this.updateMuscles.bind(this));

        this.handleClick = this.handleClick.bind(this);
        this.handleSelect = this.handleSelect.bind(this);
        this.handleRemove = this.handleRemove.bind(this);
        this.handleAdd = this.handleAdd.bind(this);
        this.handleCreate = this.handleCreate.bind(this);
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        fetch('http://petrosyan.in:8080/v1/exercise/find', {
            method: 'POST',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({input: this.state.usedMuscles}),
            credentials: 'include'
        }).then((response) => response.json())
            .then((responseJson) => {
                this.setState({ exercises: responseJson });
            })
            .catch((error) => {
                console.error(error);
            });
        fetch('http://petrosyan.in:8080/v1/user/templates/workout', {
            method: 'GET',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
            },
            credentials: 'include'
        }).then((response) => response.json())
            .then((responseJson) => {
                if(responseJson != null && typeof responseJson.templates != 'undefined') {
                    console.log(responseJson.templates);
                    this.setState({ templates: responseJson.templates });
                } else{

                    this.setState({ templates: [] });
                }
            })
            .catch((error) => {
                console.error(error);
            });

    }


    updateMuscles(str) {
        if(str == null || str === "") {
            this.setState({ usedMuscles: [] });
        } else {

            this.setState({ usedMuscles: str.split(',').filter(x => x != null).map(x => parseInt(x)) });
        }
        fetch('http://petrosyan.in:8080/v1/exercise/find', {
            method: 'POST',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({input: this.state.usedMuscles}),
            credentials: 'include'
        }).then((response) => response.json())
            .then((responseJson) => {
                if(responseJson != null && typeof responseJson.error == 'undefined') {
                    this.setState({ exercises: responseJson });
                } else {
                    this.setState({ exercises: [] });
                }
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

    handleChange(event) {
        this.setState({value: event.target.value});
    }
    handleSubmit(event) {
        event.preventDefault();
    }
    handleSelect(id) {
        if(id == this.state.currentActiveTemplate) {
            this.setState({ currentActiveTemplate: -1 });
        } else {
            this.setState({ currentActiveTemplate: id });
        }

    }

    handleCreate() {
        fetch('http://petrosyan.in:8080/v1/workout/create', {
            method: 'GET',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
            },
            credentials: 'include'
        }).then((response) => response.json())
            .then((responseJson) => {
                if(responseJson != null && typeof responseJson.error == 'undefined') {
                    var templates = this.state.templates;
                    responseJson.WorkoutExercises = [];
                    templates.push(responseJson);
                    this.setState({ templates: templates });
                }
            })
            .catch((error) => {
                console.error(error);
            });
    }

    handleAdd(id) { // TODO
        if(-1 == this.state.currentActiveTemplate) {
            alert("Please, select template!");
        } else {
            fetch('http://petrosyan.in:8080/v1/workout/'+this.state.currentActiveTemplate+'/add/'+id, {
                method: 'GET',
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/json',
                },
                credentials: 'include'
            }).then((response) => response.json())
                .then((responseJson) => {
                    if(responseJson != null && typeof responseJson.error == 'undefined') {
                        var templates = this.state.templates;
                        for(var i=0; i < templates.length; i++) {
                            if (templates[i].ID == this.state.currentActiveTemplate) {
                                templates[i].WorkoutExercises.push(responseJson);
                                break;
                            }
                        }
                        this.setState({ templates: templates });
                    }
                })
                .catch((error) => {
                    console.error(error);
                });
        }
    }

    handleRemove(id, from) {
        fetch('http://petrosyan.in:8080/v1/workout/'+from+'/remove/'+id, {
            method: 'GET',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
            },
            credentials: 'include'
        }).then((response) => response.json())
            .then((responseJson) => {
                if(responseJson != null && typeof responseJson.error == 'undefined') {
                    var templates = this.state.templates;
                    for(var i=0; i < templates.length; i++) {
                        if (templates[i].ID == from) {
                            for(var j=0; j < templates[i].WorkoutExercises.length; j++) {
                                if(templates[i].WorkoutExercises[j].ID == id) {
                                    templates[i].WorkoutExercises.splice(j, 1);
                                    //break;
                                }
                            }
                            break;
                        }
                    }
                    this.setState({ templates: templates });
                }
            })
            .catch((error) => {
                console.error(error);
            });
    }

    render(){
        return (<div>
            <Navbar/>
                <div className='profileContainer' style={{"display": "none"}}>
                    <div>
                        <img className='googlePhoto' src={this.state.profile.Image}/>
                        <p>{this.state.profile.Name}</p>
                    </div>
                    <div className='Info'>
                        <label>Age</label>
                        <input className="effect-7" type="text" placeholder=""/>

                            <span className="focus-border" />
                        <label>Weight</label>
                            <input className="effect-7" type="text" placeholder=""/>

                                <span className="focus-border" />
                        <label>Height</label>
                                <input className="effect-7" type="text" placeholder=""/>

                                    <span className="focus-border"/>
                        <label>First Name</label>
                                    <input className="effect-7" type="text" placeholder=""/>

                                        <span className="focus-border"/>
                        <a>Choose time to train:</a>
                        <ol>{this.state.templates.map(template =>
                            <li>{template.Name} </li>
                        )}
                        </ol>

                    </div>
                </div>
            <div
                style={{"position": "absolute", left: "36vw", top: "16.7vh"}}
                >
                <Unity
                    src="/Build/build.json"
                    loader="/Build/UnityLoader.js"
                    width='28vw' height='80vh'
                />
            </div>
            <div className='exerciseListBox'>
                <div ><a className='titleForBox'>Training exercises</a></div>
                <ul style={{"padding-top": "28px"}}>

                    {this.state.exercises.map(exercise=>
                        <li key={exercise.ID}>
                            <div>
                                {<section id={exercise.ID}>
                                    <img className='imageExercise'src={exercise.Image} />
                                    <a className='titleBox'> {exercise.Title}</a>
                                    <a className='addExercise' href='#'><img src={require('../img/plus.png')} onClick={() => this.handleAdd(exercise.ID)}/></a>
                                    <a className='showDetails' href='#' onClick={() => { this.handleClick(exercise.ID) }}><img src={require('../img/menu-down.png')} style={this.state.currentShow == exercise.ID ? {transform: "rotate(180deg)"}: null}/></a>
                                </section>}
                                {this.state.currentShow == exercise.ID?
                                <div className='details'>
                                    <p>{exercise.Description}</p>
                                    <p style={{'color':'#22FF09' }}>Muscle trained:</p>
                                    <ul className='musclesList' >
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
            <div className='templateListBox'>
                <div><a className='titleForBox'>Training templates</a> <a className='getNewTemplate' href="#" onClick={() => { this.handleCreate() }}><img src={require('../img/plus.png')}/></a></div>
                <ul  className='exercisesList_ForTemplate'>
                   <ul style={{"padding-top": "28px"}}> {this.state.templates.map(template =>
                       <ul style={{"padding-top": "28px"}} onClick={() => { this.handleSelect(template.ID) }}>
                            <p style={{"color": template.ID == this.state.currentActiveTemplate?"#F7FF00":""}}>{template.Name}</p>
                            {template.WorkoutExercises.map(exercise=>
                            <li>
                                <img className='imageExercise' src={exercise.Image}/>
                                <a className='titleBox'> {exercise.Title}</a>
                                <a className='addExercise' href='#'><img src={require('../img/minus.png')} onClick={() => { this.handleRemove(exercise.ID, template.ID); }}/></a>
                            </li>
                        )}
                        </ul>
                   )}
                   </ul>
                </ul>
            </div>
            </div>
            );
    }
}

export default MainPart;

