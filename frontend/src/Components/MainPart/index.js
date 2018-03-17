import React, { Component } from 'react';
import './index.css'
import ExerciseList from '../ExerciseList';
import exercise from '../api/exercise'


class MainPart extends Component {
    render(){
        return (<div className='exerciseListBox'>
                <ExerciseList exercise ={exercise}/>
            </div>
            );
    }
}

export default MainPart;
