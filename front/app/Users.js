import React from 'react';
import {Link} from 'react-router';


const users = React.createClass({
    render(){
        return (
            <div>
                Users
                <br/>
                <Link to="/">Home</Link>
                &nbsp;
                <Link to="/about_us">About Us</Link>
                <br/>
                {this.props.children}
            </div>)
    }
});

const signIn = React.createClass({
    render(){
        return (
            <div>
                Sign In
            </div>)
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

const resetPassword = React.createClass({
    render(){
        return (
            <div>
                Reset Passwrod
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
    ResetPassword: resetPassword,
    Profile: profile
};