module Admin
  class BatmanController < ApplicationController

    def index
      render nothing: true, layout: 'admin/batman'
    end

  end
end
