import React, { Component } from 'react';
import './index.css'

class Navbar extends Component {
    render() {
        return (
            <div className='navbar'>
                <a><img className='profileIcon' src={require('../img/account.png')}/></a>
                <h2>Physical training</h2>
                <a><img className='googleicon' src={require('../img/google.png')}/></a>
            </div>
        );
    }
}

export default Navbar;
