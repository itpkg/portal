require('./custom.css');

import React from 'react'
import {Link} from 'react-router'
import {Navbar, Nav, NavItem, NavDropdown, MenuItem} from 'react-bootstrap'
import i18next from 'i18next/lib';

import Reflux from 'reflux';
import ReactMixin from 'react-mixin';

import {Actions, Store} from './flux'

const header = React.createClass({
    personalBar: function(){
        var user = this.state.current_user;
        if(user){
            return (<NavDropdown eventKey={3} title={i18next.t("users.titles.welcome", user.name)} id="basic-nav-dropdown">
                <MenuItem eventKey={3.1}>Action</MenuItem>
                <MenuItem eventKey={3.2}>Another action</MenuItem>
                <MenuItem eventKey={3.3}>Something else here</MenuItem>
                <MenuItem divider/>
                <MenuItem eventKey={3.3}>Separated link</MenuItem>
            </NavDropdown>)
        }else{
            return (<NavDropdown eventKey={3} title={i18next.t("users.titles.sign_in_or_up")} id="basic-nav-dropdown">
                <MenuItem eventKey={3.1} href="/#/users/sign-in">{i18next.t("users.titles.sign_in")}</MenuItem>
                <MenuItem eventKey={3.2} href="/#/users/sign-up">{i18next.t("users.titles.sign_up")}</MenuItem>
                <MenuItem eventKey={3.3} href="/#/users/forgot-password">{i18next.t("users.titles.forgot_your_password")}</MenuItem>
                <MenuItem eventKey={3.3} href="/#/users/confirm">{i18next.t("users.titles.did_not_receive_confirmation_instructions")}</MenuItem>
                <MenuItem eventKey={3.3} href="/#/users/unlock">{i18next.t("users.titles.did_not_receive_unlock_instructions")}</MenuItem>
            </NavDropdown>)
        }
    },
    render(){

        return (
            <Navbar inverse fixedTop fluid>
                <Navbar.Header>
                    <Navbar.Brand>
                        <a href="/#/">{this.state.site_info.title}</a>
                    </Navbar.Brand>
                    <Navbar.Toggle />
                </Navbar.Header>
                <Navbar.Collapse>
                    <Nav>
                        <NavItem eventKey={1} href="/#/">{i18next.t("links.home")}</NavItem>
                        <NavItem eventKey={2} href="/#/about-us">{i18next.t("links.about_us")}</NavItem>
                        {this.personalBar()}
                    </Nav>
                    <Nav pullRight>
                        <NavItem eventKey={1} href="/?locale=en-US">{i18next.t("links.en_us")}</NavItem>
                        <NavItem eventKey={2} href="/?locale=zh-CN">{i18next.t("links.zh_cn")}</NavItem>
                    </Nav>
                </Navbar.Collapse>
            </Navbar>
        )
    }
});

const footer = React.createClass({
    render(){
        return (<p>footer</p>)
    }
});

const noMatch = React.createClass({
    render(){
        return (<div>No match</div>)
    }
});


const version = 'v20151212';

ReactMixin.onClass(header, Reflux.connect(Store));
ReactMixin.onClass(footer, Reflux.connect(Store));

module.exports = {
    Header: header,
    Footer: footer,
    NoMatch: noMatch,
    Version: version
};