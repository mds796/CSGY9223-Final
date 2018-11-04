import { html } from '@polymer/lit-element';

import { PageViewElement } from './page-view-element.js';
import { SharedStyles } from './shared-styles.js';

/**
 * @customElement
 * @polymer
 */
class RegisterView extends PageViewElement {
  render() {
    return html`
      ${SharedStyles}
      <h2>Hello ${this.prop1}!</h2>
    `;
  }
  static get properties() {
    return {
      prop1: {type: String}
    };
  }
  constructor() {
    super();
    this.prop1 = 'Register View';
  }
}

window.customElements.define('register-view', RegisterView);
