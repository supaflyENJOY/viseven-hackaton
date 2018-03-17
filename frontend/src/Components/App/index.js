import React, { Component } from 'react';
import './index.css'
import RegistrationForm from '../RegistrationForm'
import MainPart from '../MainPart'
import 'bootstrap/dist/css/bootstrap.css'


class App extends Component {
    constructor() {
        super();
        this.state = {
            login: true
        };
    }

    render(){
        return (<div>
                {this.props.login?<RegistrationForm/>:<MainPart/>}
            </div>

        );
    }
}

export default App;
