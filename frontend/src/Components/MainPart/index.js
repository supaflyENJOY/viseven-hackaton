import React, { Component } from 'react';
import './index.css'
import MuscleList from '../MusclesList';
import muscles from '../api/muscles'


class MainPart extends Component {
    render(){
        return (<div>
                <MuscleList muscles ={muscles}/>
            </div>
            );
    }
}

export default MainPart;
