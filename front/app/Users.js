import React from 'react';
import {IndexLink} from 'react-router';
import i18next from 'i18next/lib';

import {Form} from './Form'

const users = React.createClass({
    render(){
        return (
            <div className="row">
                <div className="col-md-offset-2 col-md-8">
                    {this.props.children}
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
                    </ul>
                </div>
            </div>)
    }
});

const signIn = React.createClass({
    render(){
        return (
            <Form
                action="/users/sign_in"
                resource="users"
                title={i18next.t("users.titles.sign_in")}
                fields={[
            {id:"email", type:'email', focus:true, required: true},
            {id:"password", type:'password', required: true},
            {id:"remember_me", type:'checkbox'}
            ]}
            />)
    }
});

const signUp = React.createClass({
    render(){
        return (
            <div>
                Sign up
            </div>)
    }
});

const confirm = React.createClass({
    render(){
        return (
            <div>
                Confirm
            </div>)
    }
});

const unlock = React.createClass({
    render(){
        return (
            <div>
                Unlock
            </div>)
    }
});

const forgotPassword = React.createClass({
    render(){
        return (
            <div>
                Forgot password
            </div>)
    }
});

const changePassword = React.createClass({
    render(){
        return (
            <div>
                Change Passwrod
            </div>)
    }
});

const profile = React.createClass({
    render(){
        return (
            <div>
                Profile
            </div>)
    }
});

module.exports = {
    Users: users,
    SignIn: signIn,
    SignUp: signUp,
    Confirm: confirm,
    Unlock: unlock,
    ForgotPassword: forgotPassword,
    ChangePassword: changePassword,
    Profile: profile
};