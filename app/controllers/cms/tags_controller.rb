class Cms::TagsController < ApplicationController
  before_action :must_be_admin!, except: [:show]
  layout 'personal'

end
