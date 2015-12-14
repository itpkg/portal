const en_US = {
    translation: {
        users: {
            titles: {
                sign_in: "Sign in",
                sign_up: "Sign up",
                did_not_receive_confirmation_instructions: "Didn't receive confirmation instructions?",
                did_not_receive_unlock_instructions: "Didn't receive unlock instructions?",
                forgot_your_password: "Forgot your password?",
                change_your_password: 'Change your password'
            },
            buttons:{
                resend_confirmation_instructions: "Resend confirmation instructions",
                send_me_reset_password_instructions: "Send me reset password instructions",
                resend_unlock_instructions: "Resend unlock instructions"
            },
            fields: {
                username: 'Username',
                email: 'Email',
                password: 'Password',
                password_confirmation: 'Password confirmation',
                remember_me: 'Remember me'
            }
        },
        buttons:{
          submit: "Submit",
            reset: "Reset"
        },

        copyright: 'Building by {name}. {content}'
    }
};

const zh_CN = {
    translation: {
        users: {
            titles: {
                sign_in: "用户登录",
                sign_up: "新用户注册",
                did_not_receive_confirmation_instructions: "没有收到激活邮件？",
                did_not_receive_unlock_instructions: "没有收到解锁邮件？",
                forgot_your_password: "忘记密码？",
                change_your_password: '修改密码'
            },
            buttons:{
                resend_confirmation_instructions: "重新发送激活邮件",
                send_me_reset_password_instructions: "发送重置密码邮件",
                resend_unlock_instructions: "重新发送解锁邮件"
            },
            fields: {
                username: '用户名',
                email: '电子邮箱',
                password: '密码',
                password_confirmation: '重复输入',
                remember_me: '记住我'
            }
        },
        buttons:{
            submit: "提交",
            reset: "重写"
        },
        copyright: '使用{name}搭建。 {content}'
    }
};

module.exports = {
    enUS: en_US,
    zhCN: zh_CN
};