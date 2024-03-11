import {Component, Dispatch} from 'react';
import Routes from './routes';
import './App.css';
import {Action, State} from "./store/reducer";
import { AuthContext } from './AppContext';

export default class App extends Component {
  private readonly authState: State;
  private readonly authDispatch: Dispatch<Action>;

  constructor(props : any) {
    super(props);

    const {state, dispatch} = props;

    this.authState = state;
    this.authDispatch = dispatch;
  }

  render() {
    return (
      <AuthContext.Provider value={{
        state: this.authState,
        dispatch: this.authDispatch
      }}>
        <Routes />
      </AuthContext.Provider>
    );
  }
}
