import { localStorage } from "window-or-global";

export interface State {
  isLoggedIn: boolean;
  user: any;
}

export interface Action {
  type: string;
  payload: State;
}

export const initialState : State = {
  isLoggedIn: localStorage.getItem("isLoggedIn") ? JSON.parse(localStorage.getItem("isLoggedIn") as string) : false,
  user: localStorage.getItem("user") ? JSON.parse(localStorage.getItem("user") as string) : null
};

export function reducer(state: State, action: Action) : State {
  if (action.type === "LOGIN") {
    localStorage.setItem("isLoggedIn", JSON.stringify(action.payload.isLoggedIn))
    localStorage.setItem("user", JSON.stringify(action.payload.user))
    return {
      isLoggedIn: action.payload.isLoggedIn,
      user: action.payload.user
    };
  }

  if (action.type === "LOGOUT") {
    localStorage.clear()
    return {
      isLoggedIn: false,
      user: null
    };
  }

  return state;
}
