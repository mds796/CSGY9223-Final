import {html, PolymerElement} from '@polymer/polymer';

import {ViewStyle} from './view-style.js';

import './post-card.js';

/**
 * @customElement
 * @polymer
 */
class FeedView extends PolymerElement {
    static get template() {
        return html`
      <dom-if if=[[active]]>
        <template>
          ${ViewStyle}

          <section>           
            <textarea value="{{post::change}}" rows="7" cols="80">
            </textarea>
            <input type="submit" value="Post" on-click="submitPost"/>
          </section>  
          
          <dom-repeat items="[[feed]]">
            <template>
              <section>
                <post-card user="[[item.name]]" text="[[item.text]]"></post-card>
              </section>
            </template>
          </dom-repeat>
        </template>
      </dom-if>
    `;
    }

    static get properties() {
        return {
            active: {type: Boolean, value: false},
            post: {type: String, value: ""},
            feed: {type: Array, value: []}
        };
    }

    submitPost(e) {
        const provider = this;

        e.target.disabled = true;

        fetch('/post', {method: 'POST', body: provider.post, headers:{'Content-Type': 'text/plain'}})
            .then(response => {
                provider.post = "";
                e.target.disabled = false;
            })
            .catch(err => {
            console.log("Unable to post new message: ", err);
        });
    }
}

window.customElements.define('feed-view', FeedView);
