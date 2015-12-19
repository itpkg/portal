require('./main.css');

import React from 'react'
import {Link} from 'react-router'
import {Navbar, Nav, NavItem, NavDropdown, MenuItem, Alert} from 'react-bootstrap'
import i18next from 'i18next/lib';

const header = React.createClass({
    onSignOut: function () {

    },
    personalBar: function () {
        var user =null;//todo this.state.current_user;
        if (user) {
            return (
                <NavDropdown eventKey={3} title={i18next.t("users.titles.welcome", {name:user.name})}
                             id="basic-nav-dropdown">
                    <MenuItem eventKey={3.1} href="/#/personal/profile">{i18next.t("users.titles.profile")}</MenuItem>
                    <MenuItem divider/>
                    <MenuItem eventKey={3.3} onclick={this.onSignOut}>{i18next.t("users.titles.sign_out")}</MenuItem>
                </NavDropdown>)
        } else {
            return (<NavDropdown eventKey={3} title={i18next.t("users.titles.sign_in_or_up")} id="basic-nav-dropdown">
                <MenuItem eventKey={3.1} href="/#/users/sign-in">{i18next.t("users.titles.sign_in")}</MenuItem>
                <MenuItem eventKey={3.2} href="/#/users/sign-up">{i18next.t("users.titles.sign_up")}</MenuItem>
                <MenuItem eventKey={3.3}
                          href="/#/users/forgot-password">{i18next.t("users.titles.forgot_your_password")}</MenuItem>
                <MenuItem eventKey={3.4}
                          href="/#/users/confirm">{i18next.t("users.titles.did_not_receive_confirmation_instructions")}</MenuItem>
                <MenuItem eventKey={3.5}
                          href="/#/users/unlock">{i18next.t("users.titles.did_not_receive_unlock_instructions")}</MenuItem>
            </NavDropdown>)
        }
    },
    switchLang: function (lang) {
        localStorage.setItem("locale", lang);
        location.reload();
    },
    render(){
            //todo info
        return (
            <Navbar inverse fixedTop fluid>
                <Navbar.Header>
                    <Navbar.Brand>
                        <a href="/#/">TODO</a>
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
                        <NavItem eventKey={1}
                                 onClick={this.switchLang.bind(this, "en-US")}>{i18next.t("links.en_us")}</NavItem>
                        <NavItem eventKey={2}
                                 onClick={this.switchLang.bind(this, "zh-CN")}>{i18next.t("links.zh_cn")}</NavItem>
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
        return (<div className="col-md-offset-1 col-md-10">
            <br/>
            <Alert bsStyle="danger">
                <h4>{i18next.t("no_match")}</h4>
            </Alert>
        </div>)
    }
});


const version = 'v20151212';

module.exports = {
    Header: header,
    Footer: footer,
    NoMatch: noMatch,
    Version: version
};