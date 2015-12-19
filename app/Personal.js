import React from 'react';
import {IndexLink, History} from 'react-router';
import i18next from 'i18next/lib';


import {Form} from './Form'
import {Alert} from 'react-bootstrap'

const personal = React.createClass({
    render(){
        return <div>personal todo</div>
    }
});


const profile = React.createClass({
    render(){
        return (//todo
            <div>
                Profile
            </div>)
    }
});

module.exports = {
    Personal: personal,
    Profile: profile
};