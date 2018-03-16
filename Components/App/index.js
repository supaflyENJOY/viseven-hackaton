import React, { Component } from 'react';
import './index.css'
import RegistrationForm from '../RegistrationForm'
import MainPart from '../MainPart'



class App extends Component {
    constructor() {
        super();
        this.state = {
            login: true
        };
    }

    render(){
        return (<div>
                {this.props.login?<MainPart/> :<RegistrationForm/>}
            </div>

        );
    }
}

export default App;
