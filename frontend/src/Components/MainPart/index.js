import React, { Component } from 'react';
import './index.css'
import ExerciseList from '../ExerciseList';
import exercise from '../api/exercise';
import Navbar from '../Navbar';

class MainPart extends Component {
    render(){
        return (<div>
            <Navbar/>
            <div className='exerciseListBox'>
                <ExerciseList exercise ={exercise}/>
            </div>
            </div>
            );
    }
}

export default MainPart;
