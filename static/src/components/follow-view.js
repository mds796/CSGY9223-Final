import {html, PolymerElement} from "../../node_modules/@polymer/polymer/polymer-element.js";
import {ViewStyle} from './view-style.js';
import './user-card.js';

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

          <style>
            user-card {
              display: inline-block;
              cursor: pointer;
            }
          </style>

          <section>
            <label for="query">Type a search query, then press Enter to search for users:</label>
            <input id="query" type="text" value="{{query::change}}" placeholder="Username"/>
          </section>
            
            <dom-repeat items="[[follows]]" itemsIndexAs="index">
              <template>
                <section>                  
                    <label for="[[item.name]]">
                        <input type="checkbox" id="[[item.name]]" name="[[item.name]]" value="[[index]]" checked$="[[item.followed]]" on-change="toggleFollow">
                        <user-card user="[[item.name]]"></user-card>
                    </label>
                </section>
              </template>
            </dom-repeat>
        </template>
      </dom-if>
    `;
    }

    static get properties() {
        return {
            active: {
                type: Boolean,
                value: false
            },
            query: {
                type: String,
                value: "",
                observer: "fetchFollows"
            },
            follows: {
                type: Array,
                value: []
            }
        };
    }

    fetchFollows(query, _) {
        const provider = this;

        fetch('/follows?query=' + query).then(response => {
            return response.json();
        }).then(data => {
            provider.follows = data.follows;
        }).catch(err => {
            console.log("Unable to fetch follows: ", err);
        });
    }

    toggleFollow(e) {
        const index = e.target.value;
        const path = 'follows.' + index + '.followed';

        this.set(path, !this.follows[index].followed);

        const body = JSON.stringify({name: e.target.name, followed: this.follows[index].followed});

        fetch('/follow', {method: 'POST', body: body, headers:{'Content-Type': 'application/json'}}).catch(err => {
            console.log("Unable to toggle follow: ", err);
        });
    }
}

window.customElements.define('follow-view', FollowView);