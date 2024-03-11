import {Action, State} from "./store/reducer";
import {createContext, Dispatch} from "react";

export interface IAuthContext {
  state: State;
  dispatch: Dispatch<Action>
}

export const AuthContext = createContext<IAuthContext | undefined>(undefined);
