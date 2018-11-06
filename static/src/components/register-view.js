import { html, PolymerElement } from '@polymer/polymer';

import { ViewStyle } from './view-style.js';

/**
 * @customElement
 * @polymer
 */
class RegisterView extends PolymerElement {
  static get template() {
    return html`
      <dom-if if=[[active]]>
        <template>
          ${ViewStyle}
  
          <section>
            <h2>Register view!</h2>
          </section>
        </template>
      </dom-if>
    `;
  }

  static get properties() {
    return {
      active: {type: Boolean, value: false}
    };
  }
}

window.customElements.define('register-view', RegisterView);
