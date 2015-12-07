require 'portal/book'

class BooksController < ApplicationController
  layout 'cms'

  def show
    @title = params[:id]
    name = "#{Portal::Book::ROOT}/#{@title}.html"
    if File.exists?(name)
      @books = Portal::Book.names
      @body = File.read name
    else
      head :not_found
    end
  end

  def index
    @books = Portal::Book.names
  end
end