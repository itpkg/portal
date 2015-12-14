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
            fields: {
                email: 'Email',
                password: 'Password',
                remember_me: 'Remember me'
            }
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
            fields: {
                email: '电子邮箱',
                password: '密码',
                remember_me: '记住我'
            }
        },
        copyright: '使用{name}搭建。 {content}'
    }
};

module.exports = {
    enUS: en_US,
    zhCN: zh_CN
};