import React from 'react';
import {IndexLink, History} from 'react-router';
import i18next from 'i18next/lib';

import Reflux from 'reflux';
import ReactMixin from 'react-mixin';

import {Form} from './Form'
import {Alert} from 'react-bootstrap'
import {Actions, Store} from './flux'


const users = React.createClass({
    hideIfSignIn: function () {
        if (this.state.current_user) {
            return (<Alert bsStyle="danger">
                <h4>{i18next.t("users.already_sign_in")}</h4>
            </Alert>);
        } else {
            return this.props.children;
        }
    },
    render(){
        return (
            <div className="row">
                <div className="col-md-offset-1 col-md-10">
                    <br/>
                    {this.hideIfSignIn()}
                    <br/>
                    <ul>
                        <li>
                            <IndexLink to="/users/sign-in">
                                {i18next.t('users.titles.sign_in')}
                            </IndexLink>
                        </li>
                        <li>
                            <IndexLink to="/users/sign-up">
                                {i18next.t('users.titles.sign_up')}
                            </IndexLink>
                        </li>
                        <li>
                            <IndexLink to="/users/forgot-password">
                                {i18next.t('users.titles.forgot_your_password')}
                            </IndexLink>
                        </li>
                        <li>
                            <IndexLink to="/users/confirm">
                                {i18next.t('users.titles.did_not_receive_confirmation_instructions')}
                            </IndexLink>
                        </li>
                        <li>
                            <IndexLink to="/users/unlock">
                                {i18next.t('users.titles.did_not_receive_unlock_instructions')}
                            </IndexLink>
                        </li>
                        <br/>
                        <li>
                            <IndexLink to="/">
                                {i18next.t('back_to_home')}
                            </IndexLink>
                        </li>
                    </ul>
                </div>
            </div>)
    }
});

const signIn = React.createClass({
    mixins: [History],
    onSubmit(data){
        Actions.signIn(data);
        this.history.pushState(null, `/about-us`);
    },
    render(){
        return (
            <Form
                action="/users/sign_in"
                resource="users"
                submit={this.onSubmit}
                title={i18next.t("users.titles.sign_in")}
                fields={[
            {id:"email", type:'email', focus:true, required: true},
            {id:"password", type:'password', required: true},
            {id:"remember_me", type:'checkbox'}
            ]}
            />
        )
    }
});


const signUp = React.createClass({
    render(){
        return (
            <Form
                action="/users/sign_up"
                resource="users"
                title={i18next.t("users.titles.sign_up")}
                fields={[
                {id:"username", type:'text', size:6, focus:true, required: true},
            {id:"email", type:'email', required: true},
            {id:"password", type:'password', required: true},
            {id:"password_confirmation", type:'password', required: true}

            ]}
            />
        )
    }
});

const confirm = React.createClass({
    render(){
        return (<Form
                action="/users/confirm"
                resource="users"
                title={i18next.t("users.buttons.resend_confirmation_instructions")}
                fields={[
            {id:"email", type:'email', focus:true, required: true}
            ]}
            />
        )
    }
});

const unlock = React.createClass({
    render(){
        return (<Form
            action="/users/unlock"
            resource="users"
            title={i18next.t("users.buttons.resend_unlock_instructions")}
            fields={[
            {id:"email", type:'email', focus:true, required: true}
            ]}
        />)
    }
});

const forgotPassword = React.createClass({
    render(){
        return (<Form
            action="/users/forgot_password"
            resource="users"
            title={i18next.t("users.buttons.send_me_reset_password_instructions")}
            fields={[
            {id:"email", type:'email', focus:true, required: true}
            ]}
        />)
    }
});

var parse = require('url-parse');
const resetPassword = React.createClass({
    getInitialState: function () {
        var query = parse(window.location.href, true).query;
        return {token: query.token}
    },
    render(){
        return (<Form
            action="/users/reset_password"
            resource="users"
            title={i18next.t("users.buttons.send_me_reset_password_instructions")}
            fields={[
            {id:"token", type:'hidden', value:this.state.token},
            {id:"password", type:'password', required: true},
            {id:"password_confirmation", type:'password', required: true}
            ]}
        />)
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

ReactMixin.onClass(users, Reflux.connect(Store));

module.exports = {
    Users: users,
    SignIn: signIn,
    SignUp: signUp,
    Confirm: confirm,
    Unlock: unlock,
    ForgotPassword: forgotPassword,
    ResetPassword: resetPassword,
    Profile: profile
};