import { Component, useContext } from "react";
import {AuthContext, IAuthContext} from "../AppContext";
import {redirect} from "react-router";

export default abstract class Page extends Component {
  protected authState: IAuthContext | undefined;

  protected protect() {
    if (!this.authState?.state.isLoggedIn) {
      redirect('/security/login');
    }
  }

  protected constructor(props: any) {
    super(props);

    this.authState = useContext(AuthContext);
  }

  componentDidMount() {
    const { protect } = this.props as any;
    if (protect) {
      this.protect();
    }
  }
}

export abstract class ProtectedPage extends Page {
  protected constructor(props: any) {
    super({
      ...props,
      protect: true
    });
  }
}
