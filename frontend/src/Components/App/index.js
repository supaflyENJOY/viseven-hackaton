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
            login: false
        };
        fetch('http://petrosyan.in:8080/v1/user/get', {
            method: 'GET',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
            },
            credentials: 'include'
        }).then((response) => response.json())
            .then((responseJson) => {
                if(typeof responseJson.error != 'undefined' || responseJson.logged == false) {
                    this.setState({ login: false });
                } else {

                    this.setState({ login: true });
                }
            })
            .catch((error) => {
                console.error(error);
            });
    }

    render(){
        return (<div>
                {this.state.login == false?<RegistrationForm/>:<MainPart/>}
            </div>

        );
    }
}

export default App;
