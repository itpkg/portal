class Questionnaire::AnswersController < ApplicationController
  before_action :must_be_admin!

  def destroy
    Questionnaire::Answer.destroy_all uid: params[:id]
    flash[:notice] = t 'messages.success'
    redirect_to questionnaire_report_result_path(params[:report_id])
  end
end