import React from 'react';
import {IndexLink, History} from 'react-router';
import i18next from 'i18next/lib';

import Reflux from 'reflux';
import ReactMixin from 'react-mixin';

import {Form} from './Form'
import {Alert} from 'react-bootstrap'
import {Actions, Store} from './flux'


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