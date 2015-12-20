export const SIGN_IN = "SIGN IN";
export const SIGN_OUT = "SIGN OUT";

export function signIn(user) {
    console.log("action sign out");
    return {type: SIGN_IN, user: user}
}

export function signOut() {
    console.log("action sign out");
    return {type: SIGN_OUT}
}

