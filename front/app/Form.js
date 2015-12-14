require('./custom.css');

import React from 'react';
import i18next from 'i18next/lib';
import {Input} from 'react-bootstrap';

const form = React.createClass({
    render(){
        var resource = this.props.resource;
        var fields = this.props.fields.map(function (field) {
            var name = resource + "[" + field.id + "]";
            var id = resource + "_" + field.id;
            var key = 'k-' + field.id;
            var label = i18next.t(resource + ".fields." + field.id);

            if (field.required) {
                label = "* " + label;
            }
            switch (field.type) {
                case "email":
                    return <Input key={key} type="email" label={label+":"} labelClassName="col-xs-2"
                                  wrapperClassName="col-xs-6"/>;
                case "text":
                    return <Input key={key} type="text" label={label+":"} labelClassName="col-xs-2"
                                  wrapperClassName="col-xs-10"/>;
                case "password":
                    return <Input key={key} type="password" label={label+":"} labelClassName="col-xs-2"
                                  wrapperClassName="col-xs-8"/>;
                case "checkbox":
                    return <Input key={key} type="checkbox" label={label}
                                  wrapperClassName="col-xs-offset-2 col-xs-10"/>;
                default:
                    return <input key={key} id={id} name={name} type="hidden"/>;
            }
        });
        var method = this.props.method;
        if (!method) {
            method = 'post';
        }

        return (
            <fieldset>
                <legend>{this.props.title}</legend>
                <form method={method} action={this.props.action+"?locale="+i18next.language}
                      className="form-horizontal">
                    {fields}
                </form>
            </fieldset>)
    }
});

module.exports = {
    Form: form
};