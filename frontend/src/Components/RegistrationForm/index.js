import React, { Component } from 'react';
import './index.css'


class RegistrationForm extends Component {
    constructor() {
        super()
        this.handleLogin = this.handleLogin.bind(this)
    }

    handleLogin() {
        fetch('http://petrosyan.in:8080/v1/login/google', {
            method: 'GET',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
            },
            credentials: 'include'
        }).then((response) => response.json())
            .then((responseJson) => {
                window.location = responseJson.redirect_uri;
            })
            .catch((error) => {
                console.error(error);
            });
    }

    render() {
        return (
            <div id="content">
                <div class="container">
                    <div class="row row1">
                        <div class="col-md-6 col-sm-6 col-xs-6 logo">parovozik.</div>
                        <div class="col-md-6 col-sm-6 col-xs-6 google">
                            <a href="#" onClick={this.handleLogin}><img src="/img/google.png" alt="" /></a>
                        </div>
                    </div>

                    <div class="row row2">
                        <div class="col-md-6 col-sm-6 col-xs-7">
                            <h1>Health<br />Helper</h1>
                            <p>The main goal of the project
                                is to help people follow healthy way of life .
                                The app will help you pick exercises for the group of muscles
                                you want to train. </p>
                            {/*<button>*/}
                                {/*<a href="#">Continue</a>*/}
                            {/*</button>*/}
                        </div>

                        <div class="col-md-6 col-sm-6 col-xs-5">
                            <img src="/img/index.png" alt="" />
                        </div>
                    </div>

                    <div class="row row3">
                        <div class="col-md-6 col-sm-6 col-xs-6 logo"></div>
                        <div class="col-md-6 col-sm-6 col-xs-6 google"></div>
                    </div>

                </div>
            </div>
        );
    }
}

export default RegistrationForm;
