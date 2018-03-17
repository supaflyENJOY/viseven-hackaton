import React, { Component } from 'react';
import './index.css'
import RegistrationForm from '../RegistrationForm'
import MainPart from '../MainPart'
import 'bootstrap/dist/css/bootstrap.css'
import './style.css'
import './media.css'

class App extends Component {
    constructor() {
        super();
        this.state = {
            login: true
        };
    }

    render(){
        return (<div>
                {this.state.login == true?<RegistrationForm/>:<MainPart/>}
            </div>

        );
    }
}

export default App;
