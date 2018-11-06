import { html, PolymerElement } from '@polymer/polymer';

import { ViewStyle } from './view-style.js';

/**
 * @customElement
 * @polymer
 */
class FollowView extends PolymerElement {
  static get template() {
    return html`
      <dom-if if=[[active]]>
        <template>
          ${ViewStyle}

          <section>
            <h2>Follow view!</h2>
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

window.customElements.define('follow-view', FollowView);
