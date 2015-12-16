require('./custom.css');

import $ from 'jquery';
import React from 'react';
import i18next from 'i18next/lib';
import {Input, ButtonInput} from 'react-bootstrap';

const form = React.createClass({
    getInitialState: function () {
        var sfs = {};
        this.props.fields.forEach(function (field) {
            sfs[field.id] = field.value;
        });

        return {fields: sfs}
    },
    handleChange: function (e) {
        var sfs = this.state.fields;
        sfs[e.target.id] = e.target.value;
        this.setState({fields: sfs});
    },
    handleSubmit: function (e) {
        e.preventDefault();
        switch (this.props.method) {
            default:
                $.post(this.props.action, this.state.fields, function (rs) {
                    console.log(rs);
                });
        }
    },
    render: function () {
        var handleChange = this.handleChange;
        var resource = this.props.resource;
        var fields = this.props.fields.map(function (field) {
            var key = 'k-' + field.id;
            var label = i18next.t(resource + ".fields." + field.id);

            if (field.required) {
                label = "* " + label;
            }
            switch (field.type) {
                case "email":
                    return <Input id={field.id} key={key} onChange={handleChange} type="email" label={label+":"}
                                  labelClassName="col-xs-2"
                                  wrapperClassName="col-xs-6"/>;
                case "text":
                    return <Input id={field.id} key={key} onChange={handleChange} type="text" label={label+":"}
                                  labelClassName="col-xs-2"
                                  wrapperClassName="col-xs-10"/>;
                case "password":
                    return <Input id={field.id} key={key} onChange={handleChange} type="password" label={label+":"}
                                  labelClassName="col-xs-2"
                                  wrapperClassName="col-xs-8"/>;
                case "checkbox":
                    return <Input id={field.id} key={key} onChange={handleChange} type="checkbox" label={label}
                                  wrapperClassName="col-xs-offset-2 col-xs-10"/>;
                default:
                    return <input key={key} id={id} type="hidden"/>;
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
                    <div className="form-group">
                        <div className="col-xs-offset-2 col-xs-10">
                            <button type="submit" onClick={this.handleSubmit}
                                    className="btn btn-primary">{i18next.t("buttons.submit")}</button>
                            &nbsp; &nbsp;
                            <button type="reset" className="btn btn-default">{i18next.t("buttons.reset")}</button>
                        </div>
                    </div>
                </form>
            </fieldset>)
    }
});

module.exports = {
    Form: form
};